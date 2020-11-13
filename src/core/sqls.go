package core

var AllDefaultTables = map[string]string{
	"data_sources": DataSourceTable,
	"is_installed": InstalledTable,
	"steps":        StepTable,
}

/*
CREATE TABLE `data_sources` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `title` varchar(255) NOT NULL COMMENT '名称',
  `desc` varchar(255) NOT NULL COMMENT '描述',
  `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '数据源类型,1.link,2.mysql,3.redis,4.file,5.step',
  `scheme` varchar(255) NOT NULL DEFAULT '' COMMENT '连接信息',
  `struct` varchar(255) NOT NULL DEFAULT '' COMMENT '结构信息',
  `active` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否正常，1正常0删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据源';
*/
var DataSourceTable = "CREATE TABLE `data_sources` (`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',`title` varchar(255) NOT NULL COMMENT '名称', `desc` varchar(255) NOT NULL COMMENT '描述', `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '数据源类型,1.link,2.mysql,3.redis,4.file,5.step', `scheme` varchar(255) NOT NULL DEFAULT '' COMMENT '连接信息', `struct` varchar(255) NOT NULL DEFAULT '' COMMENT '结构信息', `active` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否正常，1正常0删除', `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据源';"

/*
CREATE TABLE `is_installed` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='是否安装表';
*/
var InstalledTable = "CREATE TABLE `is_installed` ( `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID', `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='是否安装表';"

//安装完毕sql
var InstalledSql = "INSERT INTO `is_installed` (id) VALUE (NULL);"

//是否安装sql
var IsInstalledSql = "SELECT * FROM `is_installed`;"

/*
CREATE TABLE `steps` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `data_source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据源ID',
  `view_type` varchar(32) NOT NULL DEFAULT '' COMMENT '支持的view类型，英文逗号分隔，1.list,2.form,3.graph',
  `target_source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变动目标数据源ID',
  `action_type` varchar(32) NOT NULL DEFAULT '' COMMENT '支持的操作类型，英文逗号分隔，1.insert,2.update,3.delete,4.download',
  `active` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否正常，1正常0删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='步骤表';
*/
var StepTable = "CREATE TABLE `steps` ( `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID', `data_source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据源ID', `view_type` varchar(32) NOT NULL DEFAULT '' COMMENT '支持的view类型，英文逗号分隔，1.list,2.form,3.graph', `target_source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变动目标数据源ID', `action_type` varchar(32) NOT NULL DEFAULT '' COMMENT '支持的操作类型，英文逗号分隔，1.insert,2.update,3.delete,4.download', `active` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否正常，1正常0删除', `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新时间', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='步骤表';"
