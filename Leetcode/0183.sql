# Write your MySQL query statement below

# 1
select name as Customers
from Customers left outer join Orders
on Orders.customerId = Customers.id
where Orders.id is NULL;

# 2
select name as Customers
from Customers as c
where c.id not in (
    select distinct customerId
    from Orders
);