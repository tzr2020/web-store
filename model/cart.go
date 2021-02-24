package model

import (
	"log"
	"web-store/util"
)

type Cart struct {
	CartID      string
	UserID      int
	TotalCount  int         // 购物项数，通过GetTotalCount()计算得到
	TotalAmount float64     // 购物车总计金额，通过GetTotalAmount()计算得到
	CartItems   []*CartItem // 购物项
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
