package model

import (
	"log"
	"web-store/util"
)

type CartItem struct {
	CartItemID int
	CartID     string
	Product    *Product // 产品
	Count      int      // 产品数量
	Amount     float64  // 产品金额小计，通过GetAmount()计算得到
}

// GetAmount 通过购物项里的产品数量、产品里的价格相乘得到
func (cartItem *CartItem) GetAmount() float64 {
	return float64(cartItem.Count) * cartItem.Product.Price
}

// AddCartItem 数据库新增购物项
func AddCartItem(cartItem *CartItem) error {
	query := "insert into cart_items (cart_id, product_id, count, amount) values (?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Printf("准备SQL语句错误: %v", err)
		return err
	}

	_, err = stmt.Exec(cartItem.CartID, cartItem.Product.ID, cartItem.Count, cartItem.GetAmount())
	if err != nil {
		log.Printf("执行SQL语句错误: %v", err)
		return err
	}

	return nil
}

// GetCartItemByCartID 根据购物车id，从数据库获取购物项，维护购物项结构
func GetCartItemByCartID(cid string) ([]*CartItem, error) {
	query := "select id, cart_id, product_id, count, amount from cart_items"
	query += " where cart_id = ?"

	rows, err := util.Db.Query(query, cid)
	if err != nil {
		log.Printf("数据库查询购物项发生错误: %v", err)
		return nil, err
	}

	var cartItems []*CartItem
	var pid string

	for rows.Next() {
		cItem := &CartItem{}

		err = rows.Scan(&cItem.CartItemID, &cItem.CartID, &pid, &cItem.Count, &cItem.Amount)
		if err != nil {
			log.Printf("数据库扫描购物项发生错误: %v", err)
			return nil, err
		}

		// 将数据库查询的产品设置到购物项结构的产品字段
		product, err := GetProduct(pid)
		if err != nil {
			log.Printf("从数据库获取产品发生错误: %v", err)
			return nil, err
		}
		cItem.Product = product

		cartItems = append(cartItems, cItem)
	}

	return cartItems, nil
}

// GetCartItemByCartIDAndProductID 根据购物车id和产品id，从数据库获取购物项，维护购物项结构
func GetCartItemByCartIDAndProductID(cid string, pid string) (*CartItem, error) {
	query := "select id, cart_id, count, amount from cart_items"
	query += " where cart_id = ? and product_id = ?"

	cItem := &CartItem{}

	err := util.Db.QueryRow(query, cid, pid).Scan(&cItem.CartItemID, &cItem.CartID,
		&cItem.Count, &cItem.Amount)
	if err != nil {
		log.Printf("数据库扫描购物项发生错误: %v", err)
		return nil, err
	}

	// 将数据库查询的产品设置到购物项结构的产品字段
	product, err := GetProduct(pid)
	if err != nil {
		log.Printf("从数据库获取产品发生错误: %v", err)
		return nil, err
	}
	cItem.Product = product

	return cItem, nil
}

// UpdateProductCountOfCartItem 数据库更新购物项的产品数量以及金额小计
func UpdateProductCountOfCartItem(cItem *CartItem) error {
	query := "update cart_items set count = ?, amount = ?"
	query += " where cart_id = ? and Product_id = ?"

	_, err := util.Db.Exec(query, cItem.Count, cItem.GetAmount(),
		cItem.CartID, cItem.Product.ID)
	if err != nil {
		return err
	}

	return nil
}
