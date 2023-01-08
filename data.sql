INSERT INTO public.tbl_transaksi
(transaction_id, customer_id, transaction_date, harga_topup, harga_buyback, gram, saldo)
VALUES('l9z02O2Vg', 'r001', 1673148005, 80000.22, 70000.0, 1.0, 1.201);

INSERT INTO public.tbl_topup
(topup_id, customer_id, transaction_date, harga_topup, harga_buyback, gram, saldo)
VALUES('tf4d5d2Vg', 'r001', 1673145721, 80000.22, 70000.0, 0.3, 2.201);

INSERT INTO public.tbl_rekening
(customer_id, balance)
VALUES('r001', 1.2009999999999996);

INSERT INTO public.tbl_rekening
(customer_id, balance)
VALUES('r001', 1.2009999999999996);
