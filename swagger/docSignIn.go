package swagger

// swagger:operation POST /api/v1/sign-in Authorization sign-in
// ---
// summary: авторизация
// produces:
//  - application/json
// parameters:
//  - name: body
//    in: body
//    description: Параметры для входа
//    schema:
//     $ref: "#/definitions/body_credentials"
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
