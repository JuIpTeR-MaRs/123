CREATE DATABASE IF NOT EXISTS shop_db;
USE shop_db;
#用户表
CREATE TABLE users (
    uid BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) UNIQUE NOT NULL,
    nickname VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    headicon VARCHAR(255) DEFAULT NULL,
    phone VARCHAR(20) DEFAULT NULL,
    created_at VARCHAR(32),
    updated_at VARCHAR(32),
    deleted_at VARCHAR(32) DEFAULT NULL
);
#管理员表
CREATE TABLE admins (
    admin_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at VARCHAR(32),
    updated_at VARCHAR(32),
    deleted_at VARCHAR(32) DEFAULT NULL
);
#插入管理员 密码 123456
insert into admins values (1,'admin','e10adc3949ba59abbe56e057f20f883e','admin@qq.com','2025-11-05 00:00:00','2025-11-05 00:00:00',NULL);

#黑名单用户表
CREATE TABLE blacklist_uids (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    uid BIGINT NOT NULL,
    reason VARCHAR(255) DEFAULT NULL COMMENT '拉黑原因',
    created_at VARCHAR(32),
    updated_at VARCHAR(32),
    deleted_at VARCHAR(32) DEFAULT NULL
);

#商品表
CREATE TABLE products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    category_id VARCHAR(50) NOT NULL COMMENT '商品类别ID',
    description TEXT DEFAULT NULL,
    price DECIMAL(10, 2) NOT NULL,
    sku VARCHAR(100) UNIQUE NOT NULL COMMENT '商品SKU（库存单位）',
    stock INT NOT NULL COMMENT '库存数量',
    cover_image VARCHAR(255) DEFAULT NULL COMMENT '商品图片URL',
    images JSON NULL COMMENT '商品图片列表（JSON数组）',
    is_on_sale BOOLEAN DEFAULT TRUE COMMENT '是否上架',
    audit_status INT DEFAULT 0 COMMENT '审核状态：0-待审核，1-审核通过，2-审核不通过',
    audit_remark VARCHAR(255) NULL COMMENT '审核备注',
    view_count INT DEFAULT 0 COMMENT '浏览量',
    created_at VARCHAR(32),
    updated_at VARCHAR(32),
    deleted_at VARCHAR(32) DEFAULT NULL
);
#商品类别表
CREATE TABLE product_categories (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    parent_id BIGINT DEFAULT NULL COMMENT '父类别ID',
    created_at VARCHAR(32),
    updated_at VARCHAR(32),
    deleted_at VARCHAR(32) DEFAULT NULL
);

#购物车表
CREATE TABLE shopping_cart (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    user_id INT NOT NULL COMMENT '用户ID',
    product_id INT NOT NULL COMMENT '商品ID',
    quantity INT NOT NULL DEFAULT 1 COMMENT '购买数量',
    selected TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否选中',
    added_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
    UNIQUE KEY uk_user_product (user_id, product_id) -- 一个用户只能添加一次同一商品
) COMMENT='购物车表';

#订单表
CREATE TABLE order (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '订单ID',
    order_no VARCHAR(32) NOT NULL UNIQUE COMMENT '订单号',
    user_id INT NOT NULL COMMENT '用户ID',
    
    -- 商品信息（假设一单一个商品，或合并显示）
    product_id INT NOT NULL COMMENT '商品ID',
    product_name VARCHAR(255) NOT NULL COMMENT '商品名称',
    quantity INT NOT NULL DEFAULT 1 COMMENT '数量',
    unit_price DECIMAL(10, 2) NOT NULL COMMENT '单价',
    
    -- 金额
    total_amount DECIMAL(10, 2) NOT NULL COMMENT '总金额',
    pay_amount DECIMAL(10, 2) NOT NULL COMMENT '实付金额',
    
    -- 收货信息
    receiver_name VARCHAR(50) NOT NULL,
    receiver_phone VARCHAR(20) NOT NULL,
    receiver_address VARCHAR(255) NOT NULL,
    
    -- 状态
    status TINYINT NOT NULL DEFAULT 0 COMMENT '0:待付款 1:待发货 2:待收货 3:已完成 4:申请退款',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_user_status (user_id, status),
    INDEX idx_created_at (created_at)
) COMMENT='订单表';

#商品评价表
CREATE TABLE product_review (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL COMMENT '商品ID',
    user_id INT NOT NULL COMMENT '用户ID',
    
    rating TINYINT NOT NULL DEFAULT 5 COMMENT '评分 1-5',
    comment VARCHAR(500) NOT NULL COMMENT '评价内容',
    images VARCHAR(1000) DEFAULT NULL COMMENT '图片URL，多个用逗号分隔',
    
    is_anonymous BOOLEAN DEFAULT FALSE COMMENT '是否匿名',
    status TINYINT DEFAULT 1 COMMENT '1:正常 0:隐藏',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_product (product_id),
    INDEX idx_user (user_id)
) COMMENT='商品评价表';