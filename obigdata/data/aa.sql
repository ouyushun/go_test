select temp1.class,sum(temp1.degree),avg(temp1.degree)  from (SELECT  students.sno AS ssno,students.sname,students.ssex,students.sbirthday,students.class, scores.sno,scores.degree,scores.cno  FROM students LEFT JOIN scores ON students.sno =  scores.sno ) temp1 group by temp1.class