консольная игра, где будут бродить человечки. постараюсь показать наибольший стек знаний golang В проекте две вети, во второй игра реализована для телеграм.

Данный проект выполнял самостостельно, если быть точнее, то мне давали ТЗ, а как это реализовать и написать код - сам.
Никто мой код не проверял, но сам вижу что код не самый лучший. Не стал переписывать код, потому что поздно понял, о неправельности
и пришлось бы переписывать всё с начала. (часть вины плохого кода в том, что задания давали частями(и каждое задание немного противоречит другому)
и пришлось подстраиваться под тесты. Проблемы которые я могу выделить:
-не пользовался ошибками, не обрабатывал, не возвращал из функций
-имеется две глобальные переменные 
-переменну Rooms нужно было сделать мапой,
сократилось бы строк 100-150
-не везде раставил мьютексы
-почти не логировал 
-модуль телеграма добавлен без вендора

в следующих проектах буду стараться не допускать эти проблемы =)
