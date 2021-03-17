/*
 * Maintained by jemo from 2021.3.17 to now
 * Created by jemo on 2021.3.17 11:33:17
 * User
 * 用户
 */

package database

const createUser = `
  CREATE TABLE IF NOT EXISTS user (
    id INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    openid VARCHAR(28) COMMENT '微信openid',
    nickname VARCHAR(32) COMMENT '微信呢称',
    sex TINYINT COMMENT '性别，1-男性，2-女性，0-未知',
    city VARCHAR(60) COMMENT '普通用户个人资料填写的城市',
    province VARCHAR(20) COMMENT '用户个人资料填写的省份',
    country VARCHAR(60) COMMENT '国家，如中国为CN',
    headimgurl VARCHAR(150) COMMENT '用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效',
    privilege VARCHAR(100) COMMENT '用户特权信息，json 数组'
  );
`
