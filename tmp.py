
# str = "({}, 'img/strings.png', 'img/img-small/strings.png', 'string', '', 'Strings', 100, NULL, '2018-08-14 07:54:19', '2019-01-11 00:28:40', NULL),"
sql = []
f = open("demofile2.txt", "a")

for i in range(10000):
    str = "({}, 'img/strings.png', 'img/img-small/strings.png', 'string', '', 'Strings', 100, NULL, '2018-08-14 07:54:19', '2019-01-11 00:28:40', NULL),".format(i)
    sql.append(str)
    f.write(str)
f.close()
print(sql)

