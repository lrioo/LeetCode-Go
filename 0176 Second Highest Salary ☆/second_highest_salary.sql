Create table If Not Exists Employee
(
    Id     int,
    Salary int
);
Truncate table Employee;

insert into Employee (Id, Salary)
values ('1', '100');

insert into Employee (Id, Salary)
values ('2', '200');

insert into Employee (Id, Salary)
values ('3', '300');


-- Write your MySQL query statement below
SELECT Salary AS SecondHighestSalary
FROM Employee
ORDER BY Salary
LIMIT 1 OFFSET 1;


/*
Write a SQL query to get the second highest salary from the Employee table.
编写一个SQL查询，获取Employee表中第二高的薪水（Salary）。

  +----+--------+
  | Id | Salary |
  +----+--------+
  | 1  | 100    |
  | 2  | 200    |
  | 3  | 300    |
  +----+--------+


For example, given the above Employee table, the query should return 200 as the second highest salary.
If there is no second highest salary, then the query should return null.
例如上述Employee表，SQL查询应该返回200作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回null。

  +---------------------+
  | SecondHighestSalary |
  +---------------------+
  | 200                 |
  +---------------------+
*/
