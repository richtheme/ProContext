# Test Task for Golang developer in ProContext

Напишите на любом языке программирования следующую задачу:

1. Вытащить из апи Центробанка (пример http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=11/11/2020) данные по переводу различных валют в рубли за последние 90 дней.
2. Результатом работы программы:  
 - нужно указать значение максимального курса валюты, название этой валюты и дату этого максимального значения.
 - нужно указать значение минимального курса валюты, название этой валюты и дату этого минимального значения.
 - нужно указать среднее значение курса рубля за весь период по всем валютам.

К файлу с программой, нужно прикрепить инструкцию по запуску.

## How to run program

Prerequisites:
```
git clone https://github.com/richtheme/ProContext
```

1 Option (using Make):
```
make run
```

2 Option:
```
go mod tidy
go run main.go
```

### Expected outpuf after ```make run```

```
Program started: Please wait..

Max currency rate: SDR, Value: 111.114000, Date: 10/04/2023
Min currency rate: Vietnam Dong, Value: 0.002907, Date: 24/01/2023
Average rates:
Singapore Dollar: 55.520743
Turkish Lira: 3.907597
Swedish Krona: 7.100318
Danish Krone: 10.689841
Indian Rupee: 0.899185
Kazakhstan Tenge: 0.163242
UAE Dirham: 18.691459
Qatari Riyal: 18.860221
S.African Rand: 4.140356
Belarussian Ruble: 26.608726
Hungarian Forint: 0.205951
Vietnam Dong: 0.002907
Moldova Lei: 3.947848
Romanian Leu: 16.171589
Thai Baht: 2.019769
British Pound Sterling: 90.158479
Egyptian Pound: 2.246436
China Yuan: 10.787460
SDR: 99.167783
Serbian Dinar: 0.629729
Australian Dollar: 50.492652
Azerbaijan Manat: 43.497272
Norwegian Krone: 7.185565
Czech Koruna: 3.354166
Swiss Franc: 80.170383
South Korean Won: 0.057683
Bulgarian lev: 40.680433
Brazil Real: 14.300383
US Dollar: 73.945358
Canadian Dollar: 54.692818
New Turkmenistan Manat: 21.127248
Ukrainian Hryvnia: 2.003980
Armenia Dram: 0.188809
Georgia Lari: 26.280940
Indonesian Rupiah: 0.004528
Hong Kong Dollar: 9.444595
Polish Zloty: 16.893346
Uzbekistan Sum: 0.006512
Tajikistan Ruble: 6.948613
Japanese Yen: 0.558748
Euro: 79.571313
Kyrgyzstan Som: 0.850348
New Zealand Dollar: 43.194684
```