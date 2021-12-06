CONNECT DATABASE northwind;
CREATE OR REPLACE WRAPPER JDBC bv_order_details_orders
    FOLDER = '/02_base_view'
    DATASOURCENAME=ds_northwind
    SQLSENTENCE='SELECT
  a.order_id
  , a.product_id
  , a.unit_price
  , a.quantity
  , a.discount
  , b.customer_id
  , b.employee_id
  , b.order_date
  , b.required_date
  , b.shipped_date
  , b.ship_via
  , b.freight
  , b.ship_name
  , b.ship_address
  , b.ship_city
  , b.ship_region
  , b.ship_postal_code
  , b.ship_country
FROM order_details AS a
JOIN orders AS b
  ON a.order_id = b.order_Id'
    OUTPUTSCHEMA (
        order_id = 'order_id' :'java.lang.Short' (sourcetypedecimals='0', sourcetypesize='5', sourcetypeid='5', sourcetypename='int2')  NOT NULL SORTABLE NOT UPDATEABLE,
        product_id = 'product_id' :'java.lang.Short' (sourcetypedecimals='0', sourcetypesize='5', sourcetypeid='5', sourcetypename='int2')  NOT NULL SORTABLE NOT UPDATEABLE,
        unit_price = 'unit_price' :'java.lang.Double' (sourcetypedecimals='8', sourcetypesize='8', sourcetypeid='7', sourcetypename='float4')  NOT NULL SORTABLE NOT UPDATEABLE,
        quantity = 'quantity' :'java.lang.Short' (sourcetypedecimals='0', sourcetypesize='5', sourcetypeid='5', sourcetypename='int2')  NOT NULL SORTABLE NOT UPDATEABLE,
        discount = 'discount' :'java.lang.Double' (sourcetypedecimals='8', sourcetypesize='8', sourcetypeid='7', sourcetypename='float4')  NOT NULL SORTABLE NOT UPDATEABLE,
        customer_id = 'customer_id' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='2147483647', sourcetypeid='1', sourcetypename='bpchar')  SORTABLE NOT UPDATEABLE,
        employee_id = 'employee_id' :'java.lang.Short' (sourcetypedecimals='0', sourcetypesize='5', sourcetypeid='5', sourcetypename='int2')  SORTABLE NOT UPDATEABLE,
        order_date = 'order_date' :'java.time.LocalDate' (sourcetypedecimals='0', sourcetypesize='13', sourcetypeid='91', sourcetypename='date')  SORTABLE NOT UPDATEABLE,
        required_date = 'required_date' :'java.time.LocalDate' (sourcetypedecimals='0', sourcetypesize='13', sourcetypeid='91', sourcetypename='date')  SORTABLE NOT UPDATEABLE,
        shipped_date = 'shipped_date' :'java.time.LocalDate' (sourcetypedecimals='0', sourcetypesize='13', sourcetypeid='91', sourcetypename='date')  SORTABLE NOT UPDATEABLE,
        ship_via = 'ship_via' :'java.lang.Short' (sourcetypedecimals='0', sourcetypesize='5', sourcetypeid='5', sourcetypename='int2')  SORTABLE NOT UPDATEABLE,
        freight = 'freight' :'java.lang.Double' (sourcetypedecimals='8', sourcetypesize='8', sourcetypeid='7', sourcetypename='float4')  SORTABLE NOT UPDATEABLE,
        ship_name = 'ship_name' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='40', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE,
        ship_address = 'ship_address' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='60', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE,
        ship_city = 'ship_city' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='15', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE,
        ship_region = 'ship_region' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='15', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE,
        ship_postal_code = 'ship_postal_code' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='10', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE,
        ship_country = 'ship_country' :'java.lang.String' (sourcetypedecimals='0', sourcetypesize='15', sourcetypeid='12', sourcetypename='varchar')  SORTABLE NOT UPDATEABLE
    );

CREATE OR REPLACE TABLE bv_order_details_orders I18N us_pst (
        order_id:int (notnull, sourcetypeid = '5', sourcetypedecimals = '0', sourcetypesize = '5'),
        product_id:int (notnull, sourcetypeid = '5', sourcetypedecimals = '0', sourcetypesize = '5'),
        unit_price:double (notnull, sourcetypeid = '7', sourcetypedecimals = '8', sourcetypesize = '8'),
        quantity:int (notnull, sourcetypeid = '5', sourcetypedecimals = '0', sourcetypesize = '5'),
        discount:double (notnull, sourcetypeid = '7', sourcetypedecimals = '8', sourcetypesize = '8'),
        customer_id:text (sourcetypeid = '1', sourcetypedecimals = '0', sourcetypesize = '2147483647'),
        employee_id:int (sourcetypeid = '5', sourcetypedecimals = '0', sourcetypesize = '5'),
        order_date:localdate (sourcetypeid = '91', sourcetypedecimals = '0', sourcetypesize = '13'),
        required_date:localdate (sourcetypeid = '91', sourcetypedecimals = '0', sourcetypesize = '13'),
        shipped_date:localdate (sourcetypeid = '91', sourcetypedecimals = '0', sourcetypesize = '13'),
        ship_via:int (sourcetypeid = '5', sourcetypedecimals = '0', sourcetypesize = '5'),
        freight:double (sourcetypeid = '7', sourcetypedecimals = '8', sourcetypesize = '8'),
        ship_name:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '40'),
        ship_address:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '60'),
        ship_city:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '15'),
        ship_region:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '15'),
        ship_postal_code:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '10'),
        ship_country:text (sourcetypeid = '12', sourcetypedecimals = '0', sourcetypesize = '15')
    )
    FOLDER = '/02_base_view'
    CACHE OFF
    TIMETOLIVEINCACHE DEFAULT
    ADD SEARCHMETHOD bv_order_details_orders(
        I18N us_pst
        CONSTRAINTS (
             ADD order_id NOS ZERO ()
             ADD product_id NOS ZERO ()
             ADD unit_price NOS ZERO ()
             ADD quantity NOS ZERO ()
             ADD discount NOS ZERO ()
             ADD customer_id NOS ZERO ()
             ADD employee_id NOS ZERO ()
             ADD order_date NOS ZERO ()
             ADD required_date NOS ZERO ()
             ADD shipped_date NOS ZERO ()
             ADD ship_via NOS ZERO ()
             ADD freight NOS ZERO ()
             ADD ship_name NOS ZERO ()
             ADD ship_address NOS ZERO ()
             ADD ship_city NOS ZERO ()
             ADD ship_region NOS ZERO ()
             ADD ship_postal_code NOS ZERO ()
             ADD ship_country NOS ZERO ()
        )
        OUTPUTLIST (customer_id, discount, employee_id, freight, order_date, order_id, product_id, quantity, required_date, ship_address, ship_city, ship_country, ship_name, ship_postal_code, ship_region, ship_via, shipped_date, unit_price
        )
        WRAPPER (jdbc bv_order_details_orders)
    );

