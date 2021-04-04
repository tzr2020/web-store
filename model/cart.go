package model

import (
	"log"
	"web-store/util"
)

type Cart struct {
	CartID      string      `json:"cart_id,omitempty"`
	UserID      int         `json:"user_id,omitempty,string"`
	TotalCount  int         `json:"total_count,omitempty,string"`  // 购物项数，通过GetTotalCount()计算得到
	TotalAmount float64     `json:"total_amount,omitempty,string"` // 购物车总计金额，通过GetTotalAmount()计算得到
	CartItems   []*CartItem `json:"cart_items,omitempty"`          // 购物项
}

// GetTotalCount 通过购物项计算得到购物项数
func (cart *Cart) GetTotalCount() int {
	return len(cart.CartItems)
}

// GetTotalAmount 通过购物项计算得到购物车总计金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}

// AddCart 数据库新增购物车，同时新增购物项
func AddCart(cart *Cart) error {
	query := "insert into carts (id, uid, total_count, total_amount) values (?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Printf("准备SQL语句错误: %v", err)
		return err
	}

	_, err = stmt.Exec(cart.CartID, cart.UserID, cart.GetTotalCount(), cart.GetTotalAmount())
	if err != nil {
		log.Printf("执行SQL语句错误: %v", err)
		return err
	}

	// 添加购物项到数据库
	for _, v := range cart.CartItems {
		err := AddCartItem(v)
		if err != nil {
			log.Println("添加购物车到数据库时，添加购物项到数据库发生错误")
			return err
		}
	}

	return nil
}

// GetCartByUserID 根据会员用户id，从数据库获取购物车，维护购物车结构的购物项字段
func GetCartByUserID(uid int) (*Cart, error) {
	query := "select id, uid, total_count, total_amount from carts"
	query += " where uid = ?"

	cart := &Cart{}

	err := util.Db.QueryRow(query, uid).Scan(&cart.CartID, &cart.UserID, &cart.TotalCount, &cart.TotalAmount)
	if err != nil {
		log.Printf("数据库扫描购物项发生错误: %v", err)
		return nil, err
	}

	cartItems, err := GetCartItemByCartID(cart.CartID)
	if err != nil {
		return nil, err
	}
	// 设置购物车结构的购物项字段
	cart.CartItems = cartItems

	return cart, nil
}

// UpdateCountAndAmountOfCart 数据库更新购物车的购物项数和总计金额
func UpdateCountAndAmountOfCart(c *Cart) error {
	query := "update carts set total_count = ?, total_amount = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, c.GetTotalCount(), c.GetTotalAmount(), c.CartID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCart 根据购物车id，从数据库删除购物车，同时删除该购物车的购物项
func DeleteCart(cartID string) error {
	query := "delete from carts where id = ?"

	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		log.Println("从数据库删除购物项发生错误:", err)
		return err
	}

	_, err = util.Db.Exec(query, cartID)
	if err != nil {
		return err
	}

	return nil
}

// GetCartsPage 查询数据库，获取购物车列表，根据当前页的页码和每页记录条数
func GetCartsPage(pageNo int, pageSize int) (carts []*Cart, err error) {
	query := `select id, uid, total_count, total_amount
		from carts limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := &Cart{}
		err = rows.Scan(&c.CartID, &c.UserID, &c.TotalCount, &c.TotalAmount)
		if err != nil {
			return nil, err
		}
		cis, err := GetCartItemByCartID(c.CartID)
		if err != nil {
			return nil, err
		}
		c.CartItems = cis
		carts = append(carts, c)
	}

	return
}

// Add 数据库添加购物车
func (c Cart) Add() (err error) {
	query := "insert into carts (id, uid, total_count, total_amount) values (?,?,?,?)"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return
	}

	c.CartID = util.CreateUUID()
	_, err = stmt.Exec(c.CartID, c.UserID, c.TotalCount, c.TotalAmount)
	if err != nil {
		return
	}

	return
}

// Update 数据库更新购物车
func (c Cart) Update() (err error) {
	query := `update carts set uid=?, total_count=?, total_amount=?	where id=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(&c.UserID, &c.TotalCount, &c.TotalAmount, &c.CartID)
	if err != nil {
		return
	}

	return
}

// Delete 数据库删除购物车
func (c Cart) Delete() (err error) {
	query := `delete from carts where id=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(c.CartID)
	if err != nil {
		return
	}

	err = DeleteCartItemByCartID(c.CartID)
	if err != nil {
		return
	}

	return
}

func GetCartByID(cid string) (*Cart, error) {
	query := "select id, uid, total_count, total_amount from carts"
	query += " where id = ?"

	cart := &Cart{}

	err := util.Db.QueryRow(query, cid).Scan(&cart.CartID, &cart.UserID, &cart.TotalCount, &cart.TotalAmount)
	if err != nil {
		log.Printf("数据库扫描购物项发生错误: %v", err)
		return nil, err
	}

	cartItems, err := GetCartItemByCartID(cart.CartID)
	if err != nil {
		return nil, err
	}
	// 设置购物车结构的购物项字段
	cart.CartItems = cartItems

	return cart, nil
}
