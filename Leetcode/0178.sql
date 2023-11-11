# Write your MySQL query statement below

# 
select score, dense_rank() over(order by score desc) as'rank'
from Scores
order by score desc;

# algorithm
select 
    s1.score,count(distinct s2.score) as 'rank'
from
    Scores as s1 inner join Scores as s2
    on s1.score<=s2.score
group by s1.score,s1.id
order by s1.score desc