-- public.tbl_harga definition

-- Drop table

-- DROP TABLE public.tbl_harga;

CREATE TABLE public.tbl_harga (
	harga_topup float8 NOT NULL,
	harga_buyback float8 NOT NULL,
	id int4 NOT NULL,
	CONSTRAINT goldprice_pk PRIMARY KEY (id)
);


-- public.tbl_rekening definition

-- Drop table

-- DROP TABLE public.tbl_rekening;

CREATE TABLE public.tbl_rekening (
	customer_id varchar NOT NULL,
	balance float8 NOT NULL,
	CONSTRAINT tbl_rekening_pk PRIMARY KEY (customer_id)
);


-- public.tbl_topup definition

-- Drop table

-- DROP TABLE public.tbl_topup;

CREATE TABLE public.tbl_topup (
	topup_id varchar NOT NULL,
	customer_id varchar NOT NULL,
	transaction_date int8 NOT NULL,
	harga_topup float4 NOT NULL,
	harga_buyback float4 NOT NULL,
	gram float4 NOT NULL,
	saldo float4 NOT NULL
);


-- public.tbl_transaksi definition

-- Drop table

-- DROP TABLE public.tbl_transaksi;

CREATE TABLE public.tbl_transaksi (
	transaction_id varchar NOT NULL,
	customer_id varchar NOT NULL,
	transaction_date int8 NOT NULL,
	harga_topup float4 NOT NULL,
	harga_buyback float4 NOT NULL,
	gram float4 NOT NULL,
	saldo float4 NOT NULL
);