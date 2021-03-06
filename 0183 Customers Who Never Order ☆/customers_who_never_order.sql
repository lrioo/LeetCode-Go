Create table If Not Exists Customers
(
    Id   int,
    Name varchar(255)
);

Create table If Not Exists Orders
(
    Id         int,
    CustomerId int
);

Truncate table Customers;
insert into Customers (Id, Name)
values ('1', 'Joe');
insert into Customers (Id, Name)
values ('2', 'Henry');
insert into Customers (Id, Name)
values ('3', 'Sam');
insert into Customers (Id, Name)
values ('4', 'Max');

Truncate table Orders;
insert into Orders (Id, CustomerId)
values ('1', '3');
insert into Orders (Id, CustomerId)
values ('2', '1');

-- Write your MySQL query statement below
SELECT Name AS Customers
FROM Customers
WHERE Id NOT IN (SELECT CustomerId FROM Orders);

SELECT Name AS Customers
FROM Customers C
WHERE NOT EXISTS(SELECT 1 FROM Orders O WHERE O.CustomerId = C.Id);

SELECT Name AS Customers
FROM Customers
         LEFT JOIN Orders ON Customers.Id = Orders.CustomerId
WHERE Orders.CustomerId IS NULL;


/*
Suppose that a website contains two tables, the Customers table and the Orders table. Write a SQL query to find all customers who never order anything.
某网站包含两个表，Customers表和Orders表。编写一个SQL查询，找出所有从不订购任何东西的客户。


Table: Customers.
Customers 表：

  +----+-------+
  | Id | Name  |
  +----+-------+
  | 1  | Joe   |
  | 2  | Henry |
  | 3  | Sam   |
  | 4  | Max   |
  +----+-------+


Table: Orders.
Orders 表：

  +----+------------+
  | Id | CustomerId |
  +----+------------+
  | 1  | 3          |
  | 2  | 1          |
  +----+------------+


Using the above tables as example, return the following:
例如给定上述表格，你的查询应返回：

  +-----------+
  | Customers |
  +-----------+
  | Henry     |
  | Max       |
  +-----------+
*/
