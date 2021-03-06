Create table If Not Exists Employee
(
    Id        int,
    Name      varchar(255),
    Salary    int,
    ManagerId int
);
Truncate table Employee;

insert into Employee (Id, Name, Salary, ManagerId)
values ('1', 'Joe', '70000', '3');

insert into Employee (Id, Name, Salary, ManagerId)
values ('2', 'Henry', '80000', '4');

insert into Employee (Id, Name, Salary, ManagerId)
values ('3', 'Sam', '60000', 'None');

insert into Employee (Id, Name, Salary, ManagerId)
values ('4', 'Max', '90000', 'None');


-- Write your MySQL query statement below

SELECT E1.Name AS Employee
FROM Employee E1,
     Employee E2
WHERE E1.ManagerId = E2.Id
  AND E1.Salary > E2.Salary;

SELECT E1.Name AS Employee
FROM Employee E1
         JOIN Employee E2 ON E1.ManagerId = E2.Id
WHERE E1.Salary > E2.Salary;


/*
The Employee table holds all employees including their managers. Every employee has an Id, and there is also a column for the manager Id.
Employee表包含所有员工，他们的经理也属于员工。每个员工都有一个Id，此外还有一列对应员工的经理的Id。

  +----+-------+--------+-----------+
  | Id | Name  | Salary | ManagerId |
  +----+-------+--------+-----------+
  | 1  | Joe   | 70000  | 3         |
  | 2  | Henry | 80000  | 4         |
  | 3  | Sam   | 60000  | NULL      |
  | 4  | Max   | 90000  | NULL      |
  +----+-------+--------+-----------+


Given the Employee table, write a SQL query that finds out employees who earn more than their managers.
For the above table, Joe is the only employee who earns more than his manager.
给定Employee表，编写一个SQL查询，该查询可以获取收入超过他们经理的员工的姓名。在上面的表格中，Joe是唯一一个收入超过他的经理的员工。

  +----------+
  | Employee |
  +----------+
  | Joe      |
  +----------+
*/
