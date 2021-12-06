CONNECT DATABASE northwind;
CREATE OR REPLACE VIEW dv_fact_order_invoice
FOLDER = '/03_derived_view'
AS
SELECT
  a.order_id
  , a.product_id
  , b.customer_id
  , b.order_date
  , b.required_date
  , b.shipped_date
  , a.unit_price * a.quantity AS total_product_amount
  , (a.unit_price * a.quantity) * a.discount AS total_discount_amount
  ,  a.unit_price * a.quantity * (1 - a.discount) AS total_amount
FROM bv_order_details AS a
JOIN bv_orders AS b
  ON a.order_id = b.order_id