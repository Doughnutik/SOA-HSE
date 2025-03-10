Сервис **UserService** отвечает за данные пользователей, СУБД **PostgreSQL**. Имеется 3 таблицы:

1) **UserInfo**: в ней хранятся общие данные о пользователе

а) *user_id* - уникальный id пользователя (первичный ключ)

б) *name* - имя

в) *surname* - фамилия

г) *login* - логин

д) *password_hash* - захешированный пароль

2) **UserPage**: в ней хранятся данные о странице пользователя

а) *user_id* - уникальный id пользователя (вторичный ключ, ссылается на user_id в UserInfo)

б) *posts_url* - список url постов на странице пользователя

в) *images_url* - список url фотографий на странице пользователя

г) *privacy_level* - уровень приватности страницы (открыта для всех/подписчиков/никого)

д) *subscribers* - список подписчиков пользователя

е) *subscribed_for* - список тех, на кого подписан пользователь

3) **UserAdditional**: допольнительная информация о пользователе

а) *user_id* - уникальный id пользователя (вторичный ключ, ссылается на user_id в UserInfo)

б) *bio* - "О себе"

в) *country* - страна

г) *city* - город

д) *date_of_birth* - дата рождения

е) *gender* - пол

ж) *occupation* - род деятельности (школьник, студент, работающий, ищущий работы и т.д.)

з) *interests* - список интересов пользователя

и) *marital_status* - личный статус (холост, в отношениях, женат/замужем и т.д.)