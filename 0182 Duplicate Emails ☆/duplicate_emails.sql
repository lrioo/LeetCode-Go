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
