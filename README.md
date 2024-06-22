# Uyga Vazifa: `Go` da `gRPC` Client Streamingni Amalga Oshirish va Ma'lumotlar Bazasi Bilan Integratsiyasi

## Maqsad
Ushbu vazifaning maqsadi `Go` da `gRPC` `client` streamingni tushunish va amalga oshirishdir. Bir nechta savdo tranzaksiyalarini serverga yuklaydigan tizim yaratadilar. Server bu tranzaksiyalarni qayta ishlaydi, ularni `PostgreSQL` ma'lumotlar bazasida saqlaydi va umumiy savdo summasi va tranzaksiyalar sonini qaytaradi.

## Vazifa Tavsifi
Siz clientdan savdo tranzaksiyalarining oqimini qabul qiladigan `gRPC` service yaratishingiz kerak bo'ladi. Server bu tranzaksiyalarni qayta ishlaydi, ularni `PostgreSQL` ma'lumotlar bazasida saqlaydi va umumiy savdo summasi va tranzaksiyalar sonini qaytaradi.

## Talablar
1. `Protobuf` Ta'rifi: `Protobuf` message va servicelarni `.proto` faylida aniqlang.
2. `gRPC Server`: Oqimdagi ma'lumotlarni qabul qilish va qayta ishlash uchun `gRPC` serverni amalga oshiring.
3. `gRPC Client`: Serverga savdo tranzaksiyalarining oqimini yuborish uchun `gRPC` clientni amalga oshiring.
4. Ma'lumotlar Bazasi Integratsiyasi: Server tranzaksiyalarni `PostgreSQL` ma'lumotlar bazasida saqlaydi.

### Ma'lumotlar Bazasi Sxemasi
`PostgreSQL` ma'lumotlar bazasini va `sales_transactions` nomli jadvalni quyidagi sxema bilan yarating:

```sql
CREATE TABLE sales_transactions (
    transaction_id VARCHAR PRIMARY KEY,
    product_id VARCHAR,
    quantity INT,
    price DECIMAL,
    timestamp BIGINT
);

```

Proto file ni mazmuni:
```proto
syntax = "proto3";

package sales;

service SalesService {
    rpc StreamSalesTransactions (stream SalesTransaction) returns (SalesSummary);
}

message SalesTransaction {
    string transaction_id = 1;
    string product_id = 2;
    int32 quantity = 3;
    float price = 4;
    int64 timestamp = 5;
}

message SalesSummary {
    float total_amount = 1;
    int32 total_transactions = 2;
}
```