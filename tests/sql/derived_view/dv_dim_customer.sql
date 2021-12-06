CONNECT DATABASE northwind;
CREATE OR REPLACE VIEW dv_dim_customer
FOLDER = '/03_derived_view'
AS
SELECT
  customer_id
  , company_name
  , contact_name
  , contact_title
  , phone
  , fax
FROM bv_customers;