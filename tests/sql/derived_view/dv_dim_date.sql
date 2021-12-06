CONNECT DATABASE northwind;
CREATE OR REPLACE VIEW dv_dim_date
FOLDER = '/03_derived_view'
AS
SELECT
  order_date AS date
  , GETYEAR(order_date) AS year
  , GETQUARTER(order_date) AS quarter_of_year
  , GETMONTH(order_date) AS month_of_year
  , GETDAY(order_date) AS day_of_month
FROM bv_orders
WHERE order_date IS NOT NULL
UNION
SELECT
  required_date AS date
  , GETYEAR(required_date) AS year
  , GETQUARTER(required_date) AS quarter_of_year
  , GETMONTH(required_date) AS month_of_year
  , GETDAY(required_date) AS day_of_month
FROM bv_orders
WHERE required_date IS NOT NULL
UNION
SELECT
  shipped_date AS date
  , GETYEAR(shipped_date) AS year
  , GETQUARTER(shipped_date) AS quarter_of_year
  , GETMONTH(shipped_date) AS month_of_year
  , GETDAY(shipped_date) AS day_of_month
FROM bv_orders
WHERE shipped_date IS NOT NULL