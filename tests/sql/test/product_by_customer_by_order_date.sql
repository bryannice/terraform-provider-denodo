SELECT
  d.product_category
  , d.product_category_description
  , d.product_name
  , c.company_name
  , b.year AS order_year
  , b.quarter_of_year AS order_quarter
  , b.month_of_year AS order_month
  , b.date AS order_date
  , a.total_product_amount
  , a.total_discount_amount
  , a.total_amount
FROM northwind.dv_fact_order_invoice AS a
JOIN northwind.dv_dim_date AS b
  ON a.order_date = b.date
JOIN northwind.dv_dim_customer AS c
  ON a.customer_id = c.customer_id
JOIN northwind.dv_dim_product AS d
  ON a.product_id = d.product_id