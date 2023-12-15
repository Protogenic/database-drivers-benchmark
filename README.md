# Benchmark
In this project I compared the work of different database drivers using four queries and a large data set of New York taxi trips.

## Setup
**Dataset** — “New York Taxi Rides”, 209.3MB, 1,048,576 rows. 

The benchmark includes four queries executed on 3 database drivers. Each query was executed 10 times on each driver and the graph shows the average values.
![Benchmark chart](https://github.com/Protogenic/database-drivers-benchmark/assets/82569672/8be7554a-aca8-4b46-810a-61bd6ba0ec06)

### Queries
**Query 1:**
```bash
SELECT cab_type, count(*) FROM trips GROUP BY 1;
```

**Query 2:**
```bash
SELECT passenger_count, avg(total_amount) 
FROM trips 
GROUP BY 1;
```

**Query 3:**
```bash
SELECT
   passenger_count, 
   extract(year from pickup_datetime),
   count(*)
FROM trips
GROUP BY 1, 2;
```

**Query 4:** 
```bash
SELECT
    passenger_count,
    extract(year from pickup_datetime),
    round(trip_distance),
    count(*)
FROM trips
GROUP BY 1, 2, 3
ORDER BY 2, 4 desc;
```

## Drivers
### Postgres
PostgreSQL is quite well optimized for general use cases. This means that read performance is often better than Sqlite and other databases under complex workloads. In particular, the speed of Postgres stands out on the first and second queries, where there is no need to use ORDER BY statement and extract specific data from rows. The main disadvantage is that the execution time begins to seriously increase on queries that require data ordering, as can be seen in the fourth query.
### Sqlite
First of all, Sqlite stands out for its ease of integration, it does not need to be installed and is easy to use. It’s also worth noting the variety of drivers that can be used with Sqlite. From the minuses, the graph shows that Sqlite suffers most of all on the same queries as Postgres, but is inferior to it in terms of execution time on all four.
### Dataframe
The first thing that catches your eye is the stability in execution time on various queries, although Dataframe is inferior in time to Postgres on the first two queries, with increasing complexity of queries the execution time on Dataframe practically does not change. Another big advantage in this project is the ability to work directly with a csv table, without complex import of data into databases.

### Comparison
Postgres requires more time and effort to set up and maintain than SQLite. SQLite may be the better choice if you need a simple and lightweight database engine that is easy to use and requires minimal setup. PostgreSQL may be the better choice if you need a more powerful and fast engine. Dataframe is more suitable for large data mining and manipulations, the query optimizations and so on.
Overall, at places, where simple data manipulations, like data retrieval, handling, join, filtering is performed, what were executed in this project, PostgreSQL can be considered the best choice to use between these three drivers.
