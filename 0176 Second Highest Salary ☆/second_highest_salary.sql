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
GROUP BY Salary
UNION ALL
(SELECT NULL AS SecondHighestSalary)
ORDER BY SecondHighestSalary DESC
LIMIT 1 OFFSET 1;

SELECT MAX(Salary) AS SecondHighestSalary
FROM Employee
WHERE Salary NOT IN (SELECT MAX(Salary) FROM Employee);

SELECT MAX(Salary) AS SecondHighestSalary
FROM Employee
WHERE Salary < (SELECT MAX(Salary) FROM Employee);

SELECT Salary AS SecondHighestSalary
FROM Employee E1
WHERE 1 = (SELECT COUNT(DISTINCT E2.Salary)
           FROM Employee E2
           WHERE E2.Salary > E1.Salary);

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
