# Write your MySQL query statement below

# join query
delete p1
from Person p1,Person p2
where p1.Email = p2.Email and p1.id > p2.id;

# sub query
delete
from Person
where id not in(
    select id
    from (
        select min(id) as id
        from Person
        group by email
    ) as m
);