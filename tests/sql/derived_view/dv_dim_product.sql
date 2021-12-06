CONNECT DATABASE northwind;
CREATE OR REPLACE VIEW dv_dim_product
FOLDER = '/03_derived_view'
AS
SELECT
  a.product_id
  , b.category_name AS product_category
  , b.description AS product_category_description
  , a.product_name
  , a.quantity_per_unit
FROM bv_products AS a
JOIN bv_categories AS b
  ON a.category_id = b.category_id