Create table If Not Exists Person
(
    Id    int,
    Email varchar(255)
);
Truncate table Person;

insert into Person (Id, Email)
values ('1', 'a@b.com');

insert into Person (Id, Email)
values ('2', 'c@d.com');

insert into Person (Id, Email)
values ('3', 'a@b.com');

-- Write your MySQL query statement below
SELECT Email
FROM Person P
GROUP BY Email
HAVING COUNT(Email)>1;

SELECT P.Email
FROM (SELECT Email, COUNT(Email) AS cnt FROM Person GROUP BY Email) P
WHERE P.cnt > 1;

SELECT DISTINCT P1.Email
FROM Person P1
         LEFT JOIN Person P2 ON P1.Email = P2.Email
WHERE P1.Id <> P2.Id;

/*
Write a SQL query to find all duplicate emails in a table named Person.
编写一个SQL查询，查找Person表中所有重复的电子邮箱。

  +----+---------+
  | Id | Email   |
  +----+---------+
  | 1  | a@b.com |
  | 2  | c@d.com |
  | 3  | a@b.com |
  +----+---------+


For example, your query should return the following for the above table:
例如，根据以上输入，你的查询应返回以下结果：

  +---------+
  | Email   |
  +---------+
  | a@b.com |
  +---------+


Note: All emails are in lowercase.
说明：所有电子邮箱都是小写字母。
*/
