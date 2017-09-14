##操作前需配置相关环境变量并导入相关库文件
##运行dbSelector程序所需配置如下：
##
	1、下载如下3个oracle客户端相关应用包并运行安装
		oracle-instantclient11.2-basic-11.2.0.1.0-1.x86_64.rpm
		oracle-instantclient11.2-devel-11.2.0.1.0-1.x86_64.rpm
		oracle-instantclient11.2-sqlplus-11.2.0.1.0-1.x86_64.rpm
		注：为防止安装冲突影响后续操作，请顺序安装（使用rpm -ivh 包名 解压）
	2、配置相关环境变量
		$export  ORACLE_HOME = /usr/lib/oracle/11.2/client64
		$export  LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/lib/oracle/11.2/client64/lib:/usr/local/lib
	3、配置TNS文件路径
		mkdir -p /usr/lib/oracle/11.2/client64/network/admin
###mysql
	前端需输入登录mysql数据库的账号密码,ip,port以及进入数据库要选择的database,如果未选择database,登陆成功后需在sql执行框中选择数据库，先查看有哪些databases(show databases),然后执行use database
	测试用例：select * from webdemo_admin
			  create table test(id int(4),name varchar(20),age varchar(20),primary key (id))
			  alter table webdemo_admin add transactor varchar(10) not Null;
			  drop table webdemo_admin

###sqlite3
	前端只需输入要操作的database即可，注意路径要写全，否则默认数据库将保存在当前工作目录，例如
	/root/Molen.db
	测试用例：查看表以及表结构： select * from sqlite_master where type="table"

###Hbase
	前端只需输入zookeeper集群的IP地址和端口以及在jsonText文本框中输入HBase对应的主机名
	jsonText文本用例：
	[
	    {
			 "HbaseIp": "172.16.13.69",
			 "HbDomain": "breath"
		 },
		 {
			 "HbaseIp": "172.168.13.68",
			 "HbDomain": "breath"
		 }
	]
	测试用例：put test,test1,f1:c1,value1
			  get test,test1,f1:c1
	          delete test,test1

###Oracle
	前端需输入登录oracle数据库的用户,密码以及要操作的database（默认为orcl）
	测试用例：select * from test
###Redis
	前端需输入redis的IP 以及 Port，目前仅支持五类数据类型：string/set集合/list链表/hash集合/sortedSet集合
	测试用例如下：
	string类型：
		增：set key value
		删: del key value
		改: set key newValue
		查: get key
	set集合类型：
		增：sadd key value
		删：srem key value
		改：先删除再增加
		查：smembers key
	list链表类型：
		增：lpush key value
		删：lrem key index value
		改：lset key index value
		查：lrange key FirstIndex LastIndex (例如查询第一个到第三个数据：lrange key 0 3)
			lrange key 0 -1 (查询所有数据)
	hash集合类型：
		增：hset key keyMember keyMemberValue 
		删：hdel key keyMember
		改：hset key keyMember keyMemberValue（重新赋值即可）
		查：hkeys key (查询集合中所有的key)
			hvals key (查询集合中所有的value)
			hgetall key (查询所有的key对应的value)
	sortedSet集合类型：
		增：zadd key num value (插入value并置其序号为num)
		删：zrem key value 
		改：先删除再增加
		查：zrange key FirstIndex LastIndex (用法同list中的lrange)
