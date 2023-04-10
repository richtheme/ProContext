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
git clone 
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
Min currency rate: China Yuan, Value: 10.001600, Date: 13/01/2023
Average rates:
Indian Rupee: 86.743724
China Yuan: 10.769076
SDR: 98.978976
Armenia Dram: 18.842679
Belarussian Ruble: 26.589197
Brazil Real: 14.267198
Danish Krone: 21.561869
British Pound Sterling: 89.970738
Moldova Lei: 39.380633
Swedish Krona: 70.876441
Vietnam Dong: 28.680696
Kazakhstan Tenge: 16.287298
New Zealand Dollar: 42.622384
Ukrainian Hryvnia: 20.004378
Azerbaijan Manat: 43.417960
Hong Kong Dollar: 88.069950
Romanian Leu: 16.139311
Kyrgyzstan Som: 84.898550
Uzbekistan Sum: 65.008872
Japanese Yen: 55.776611
Australian Dollar: 50.418639
Bulgarian lev: 40.596526
UAE Dirham: 18.444157
Canadian Dollar: 54.597157
New Turkmenistan Manat: 21.088724
Swiss Franc: 80.006800
S.African Rand: 41.359224
Hungarian Forint: 20.541557
US Dollar: 73.810528
Qatari Riyal: 18.610697
Tajikistan Ruble: 69.410302
Polish Zloty: 16.859727
Turkish Lira: 39.016202
Georgia Lari: 25.922366
Euro: 79.408759
Egyptian Pound: 22.170381
Norwegian Krone: 71.770622
Serbian Dinar: 62.129102
South Korean Won: 57.615682
Indonesian Rupiah: 44.672260
Singapore Dollar: 55.418843
Thai Baht: 19.932709
Czech Koruna: 33.463048
```