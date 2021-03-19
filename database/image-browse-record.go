/*
 * Maintained by jemo from 2021.3.19 to now
 * Created by jemo on 2021.3.19 15:35:28
 * Image Browse Record
 * 主图浏览记录
 */

package database

const createImageBrowseRecord = `
  CREATE TABLE IF NOT EXISTS imageBrowseRecord (
    userId INTEGER UNSIGNED NOT NULL,
    imageId INTEGER UNSIGNED NOT NULL,
    start DATETIME NOT NULL COMMENT '浏览开始时间',
    end DATETIME NOT NULL COMMENT '浏览结束时间',
    duration INTEGER UNSIGNED NOT NULL COMMENT '浏览时长/毫秒'
  )
`
