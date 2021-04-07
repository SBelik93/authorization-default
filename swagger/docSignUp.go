package swagger


// swagger:operation POST /api/v1/sign-up Authorization sign-up
// ---
// summary: регистрация
// produces:
//  - application/json
// parameters:
//  - name: body
//    in: body
//    description: Данные пользователя
//    schema:
//     $ref: "#/definitions/user"
// responses:
//  '200':
//   description: Временный пароль изменен
//   schema:
//    allOf:
//     - $ref: '#/definitions/http_response'
//     - type: object
//       properties:
//        result:
//         $ref: '#/definitions/user'
//  '400':
//   description: Ошибка/и изменения пароля
//   schema:
//    $ref: '#/definitions/http_response'