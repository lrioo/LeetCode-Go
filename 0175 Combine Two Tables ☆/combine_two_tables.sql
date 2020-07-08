Create table Person
(
    PersonId  int,
    FirstName varchar(255),
    LastName  varchar(255)
);
Create table Address
(
    AddressId int,
    PersonId  int,
    City      varchar(255),
    State     varchar(255)
);

Truncate table Person;
insert into Person (PersonId, LastName, FirstName)
values ('1', 'Wang', 'Allen');

Truncate table Address;
insert into Address (AddressId, PersonId, City, State)
values ('1', '2', 'New York City', 'New York');


-- Write your MySQL query statement below
SELECT P.FirstName, P.LastName, A.City, A.State
FROM Person P
         LEFT JOIN Address A ON P.PersonId = A.PersonId;


/*
Table: Person
表1: Person

  +-------------+---------+
  | Column Name | Type    |
  +-------------+---------+
  | PersonId    | int     |
  | FirstName   | varchar |
  | LastName    | varchar |
  +-------------+---------+
  PersonId is the primary key column for this table.
  PersonId 是上表主键


Table: Address
表2: Address

  +-------------+---------+
  | Column Name | Type    |
  +-------------+---------+
  | AddressId   | int     |
  | PersonId    | int     |
  | City        | varchar |
  | State       | varchar |
  +-------------+---------+
  AddressId is the primary key column for this table.
  AddressId 是上表主键


Write a SQL query for a report that provides the following information for each person in the Person table,
regardless if there is an address for each of those people:
编写一个SQL查询，满足条件：无论person是否有地址信息，都需要基于上述两表提供person的以下信息：

  FirstName, LastName, City, State
*/
