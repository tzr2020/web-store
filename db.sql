DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL UNIQUE COMMENT '用户名称',
    `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
    `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户邮箱',
    `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
    `sex` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '男' COMMENT '性别',
    `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '/img/avatar/default.jpg' COMMENT '头像',
    `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电话号',
    `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '国家',
    `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省',
	`city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '市',
	PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '用户表';

INSERT INTO `users` VALUES (1, 'user', '12345678', 'user@qq.com', 'zzz', '女', '/img/avatar/default.jpg', '13588763355', '中国', '山东', '青岛');
INSERT INTO `users` VALUES (NULL, 'user2', '12345678', 'user2@163.com', NULL, DEFAULT, DEFAULT, NULL, NULL, NULL, NULL);



DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '产品id',
    `category_id` int(11) NOT NULL COMMENT '类别id',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `price` decimal(10, 2) NOT NULL COMMENT '价格',
    `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存',
    `sales` int(11) NOT NULL DEFAULT 0 COMMENT '销量',
    `img_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'static/img/default_product.jpg' COMMENT '图片路径',
    `detail` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'static/img/default_product_detail.jpg' COMMENT '详情描述图片路径',
    `hot_point` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '热点描述',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '产品表';

INSERT INTO `products` VALUES (1, 1, '新鲜红富士苹果 水果3斤装', 12.00, 100, 10, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 1, '无籽青提 2袋装 净重1.6kg', 79.00, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 1, '赣南脐橙5kg装', 99.50, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);

INSERT INTO `products` VALUES (NULL, 2, '鲜活花螺500g', 68.00, 134, 23, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 2, '大闸蟹鲜活螃蟹现货 4对8只', 268.00, 321, 211, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 2, '新鲜海捕大黄鱼 袋装 300-400g/条*4条', 79.00, 453, 321, DEFAULT, DEFAULT, DEFAULT);

INSERT INTO `products` VALUES (NULL, 3, '山林散养 纯粮食喂养 前蹄后蹄2个装', 45.00, 234, 35, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 3, '牛腩肉 2斤装', 68.00, 124, 35, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 3, '活杀冷冻鸡翅肉 鸡翅中2kg', 128.00, 532, 342, DEFAULT, DEFAULT, DEFAULT);

INSERT INTO `products` VALUES (NULL, 4, '牛丸组合1kg', 48.00, 456, 111, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 4, '墨鱼丸 烧烤火锅食材丸子500g', 39.00, 224, 53, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 4, '1斤马鲛丸', 54.60, 432, 11, DEFAULT, DEFAULT, DEFAULT);

INSERT INTO `products` VALUES (NULL, 5, '新鲜生菜 西餐沙拉食材 1斤', 28.00, 563, 345, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 5, '新鲜山药 2500g', 70.00, 231, 86, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `products` VALUES (NULL, 5, '散养土鸡蛋 30枚', 49.90, 864, 93, DEFAULT, DEFAULT, DEFAULT);



DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
    `session_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `user_id` int(11) NOT NULL,
    PRIMARY KEY (`session_id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci;



DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '产品类别id',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `p_id` int(11) NULL DEFAULT 0 COMMENT '父级id',
    `level` int(11) NULL DEFAULT 0 COMMENT '层级',
    `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '图片路径',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '商品类别表';

INSERT INTO `categories` VALUES (1, '新鲜水果', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `categories` VALUES (2, '海鲜水产', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `categories` VALUES (3, '精选肉类', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `categories` VALUES (4, '冷饮冻食', DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `categories` VALUES (5, '蔬菜蛋品', DEFAULT, DEFAULT, DEFAULT);



DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts` (
    `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '购物车id',
    `uid` int(11) NOT NULL COMMENT '用户id',
    `total_count` int(11) NOT NULL COMMENT '购物项数',
    `total_amount` decimal(10, 2) NOT NULL COMMENT '购物车总计金额',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '购物车表';



DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '购物项id',
    `cart_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '购物车id',
    `product_id` int(11) NOT NULL COMMENT '产品id',
    `count` int(11) NOT NULL COMMENT '产品数量',
    `amount` decimal(10, 2) NOT NULL COMMENT '产品小计金额',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '购物项表';



DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
    `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单id',
    `uid` int(11) NOT NULL COMMENT '用户id',
    `total_count` int(11) NOT NULL COMMENT '订单项数',
    `total_amount` decimal(20, 2) NOT NULL COMMENT '产品金额',
    `payment` decimal(20, 2) NOT NULL COMMENT '支付金额=产品金额+运费',
    `payment_type` tinyint(2) NULL DEFAULT 1 COMMENT '支付方式：1-在线支付，2-货到付款',

    `ship_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '快递单号',
    `ship_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '快递公司',
    `ship_fee` decimal(20, 2) NOT NULL COMMENT '运费',

    `order_status` int(11) NULL DEFAULT 0 COMMENT '状态字典的状态码',
    `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建时间',
    `update_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '更新时间',
    `payment_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '支付时间',
    `ship_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '发货时间',
    `received_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '收货时间',
    `finish_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '完成时间',
    `close_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '关闭时间',
    `status` tinyint(4) NULL DEFAULT 1 COMMENT '状态：0-禁用，1-正常，-1-删除', 
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '订单表';



DROP TABLE IF EXISTS `order_payment_type`;
CREATE TABLE `order_payment_type` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `code` tinyint(10) NOT NULL UNIQUE,
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "",
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '订单支付类型表';

INSERT INTO `order_payment_type` VALUES (DEFAULT, 1, 'ALIPAY', '支付宝');
INSERT INTO `order_payment_type` VALUES (DEFAULT, 2, 'WXPAY', '微信支付');
INSERT INTO `order_payment_type` VALUES (DEFAULT, 3, 'HDFK', '货到付款');



DROP TABLE IF EXISTS `order_status`;
CREATE TABLE `order_status` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `code` tinyint(10) NOT NULL UNIQUE,
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "",
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '订单状态字典表';

INSERT INTO `order_status` VALUES (1, -1, 'CREATE_FAILED', '创建订单失败');
INSERT INTO `order_status` VALUES (2, 0, 'WAIT_BUYER_PAY', '等待买家付款');
INSERT INTO `order_status` VALUES (3, 1, 'PAYMENT_CONFIRMING', '付款确认中');
INSERT INTO `order_status` VALUES (4, 2, 'BUYER_PAYMENT_FAILED', '买家付款失败');
INSERT INTO `order_status` VALUES (5, 3, 'BUYER_PAYMENT_SUCCESS', '买家付款成功');
INSERT INTO `order_status` VALUES (6, 4, 'SELLER_DELIVEED', '卖家已发货');
INSERT INTO `order_status` VALUES (7, 5, 'BUYER_RECEIVED', '买家已收货/交易完成');
INSERT INTO `order_status` VALUES (8, 6, 'GOODS_RETURNING', '退货中');
INSERT INTO `order_status` VALUES (9, 7, 'GOODS_RETURNED_SUCCESS', '退货成功');
INSERT INTO `order_status` VALUES (10, 8, 'ORDER_CLOSED', '订单关闭');



DROP TABLE IF EXISTS `order_addresses`;
CREATE TABLE `order_addresses` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `order_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单id',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
    `tel` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '电话号',
    `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '省',
    `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '市',
    `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '区',
    `street` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '街道',
    `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '邮编',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '订单地址表';



DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `order_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单id',
    `product_id` int(11) NOT NULL COMMENT '产品id',
    `count` int(11) NOT NULL COMMENT '产品数量',
    `amount` decimal(10, 2) NOT NULL COMMENT '产品小计金额',
    PRIMARY KEY (`id`) 
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '订单项表';



DROP TABLE IF EXISTS `addresses`;
CREATE TABLE `addresses` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `uid` int(11) NOT NULL COMMENT '用户id',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
    `tel` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '电话号',
    `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '省',
    `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '市',
    `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '区',
    `street` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '街道',
    `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT "" COMMENT '邮编',
    `is_default` int(3) NULL DEFAULT 0 COMMENT '是否默认：0-否，1-是',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '用户预存的收货地址表';

INSERT INTO `addresses` VALUES (1, 1, '张三', '15837694786', '山东', '青岛', '城阳区', '春阳路11号', '262621', 1);
INSERT INTO `addresses` VALUES (2, 1, '张三', '15837694786', '山东', '青岛', '城阳区', '春阳路12号', '262621', DEFAULT);



DROP TABLE IF EXISTS `index_new_products`;
CREATE TABLE `index_new_products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_id` int(11) NOT NULL COMMENT '产品id',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '首页新品产品表';

INSERT INTO `index_new_products` VALUES (DEFAULT, 5);
INSERT INTO `index_new_products` VALUES (DEFAULT, 6);
INSERT INTO `index_new_products` VALUES (DEFAULT, 7);
INSERT INTO `index_new_products` VALUES (DEFAULT, 8);

DROP TABLE IF EXISTS `index_hot_products`;
CREATE TABLE `index_hot_products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_id` int(11) NOT NULL COMMENT '产品id',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '首页热销产品表';

INSERT INTO `index_hot_products` VALUES (DEFAULT, 1);
INSERT INTO `index_hot_products` VALUES (DEFAULT, 2);
INSERT INTO `index_hot_products` VALUES (DEFAULT, 3);
INSERT INTO `index_hot_products` VALUES (DEFAULT, 4);
INSERT INTO `index_hot_products` VALUES (DEFAULT, 5);



DROP TABLE IF EXISTS `index_recom_products`;
CREATE TABLE `index_recom_products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_id` int(11) NOT NULL COMMENT '产品id',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '首页推荐产品表';

INSERT INTO `index_recom_products` VALUES (DEFAULT, 9);
INSERT INTO `index_recom_products` VALUES (DEFAULT, 10);
INSERT INTO `index_recom_products` VALUES (DEFAULT, 11);
INSERT INTO `index_recom_products` VALUES (DEFAULT, 12);
INSERT INTO `index_recom_products` VALUES (DEFAULT, 13);



DROP TABLE IF EXISTS `nav_products`;
CREATE TABLE `nav_products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_id` int(11) NOT NULL COMMENT '产品id',
    PRIMARY KEY (`id`)
)CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT '导航栏产品表';

INSERT INTO `nav_products` VALUES (DEFAULT, 1);
INSERT INTO `nav_products` VALUES (DEFAULT, 2);
INSERT INTO `nav_products` VALUES (DEFAULT, 3);