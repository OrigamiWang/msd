INSERT `user`(`name`, `age`, `sex`)
VALUES
('origami', 21, 'male'),
('jk', 21, 'male'),
('wang', 21, 'female'),
('ming', 21, 'other'),
('dk', 21, 'male'),
('bgk', 21, 'female');

ALTER TABLE config MODIFY COLUMN conf VARCHAR(5000) NULL COMMENT '配置';

INSERT INTO svc_config(`svc_name`, `desc`, `env`, `conf`)
VALUES
('manage', 'manage开发环境的配置', 0, 'config...');

UPDATE svc_config SET conf='# 数据库配置
# 双引号作用：支持转义、字符串类型、支持特殊字符
DATABASES:
  - KEY: sample_mysql4
    TYPE: mysql
    NAME: test
    HOST: 127.0.0.1
    PORT: 3306
    USER: root
    PASSWORD: "abc123"
  - KEY: sample_redis1
    TYPE: redis
    NAME: sample_redis
    HOST: 127.0.0.1
    PORT: 6379
    USER: root
    PASSWORD: ""
ext:
  testStruct:
    Key1: vv1
    Key2: vv2
  KEY: Hello Word!
  key2: values2
  key3Int: 3
  key4Float: 3.1
  key5Bool: TRUE' WHERE id = 3;