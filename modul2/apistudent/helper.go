package main

import "github.com/gofiber/fiber/v2"

func ok(c *fiber.Ctx, message string, data any) error {  //set success true, menampilkan message dan data dari parameter webresponse
    return c.Status(fiber.StatusOK).JSON(WebResponse{ //set status = 200 get berhasil ditemukan
        Success: true, Message: message, Data: data,
    })

	// penggunaan : ok(c, "isiMessage", namaVar)
}
 
func okList(c *fiber.Ctx, message string, data any, meta *Meta) error {  //set success true, menampilkan message, data, dan meta dari parameter webresponse
    return c.Status(fiber.StatusOK).JSON(WebResponse{ //set status = 200 get berhasil ditemukan dengan meta (info tambahan penyajian data)
        Success: true, Message: message, Data: data, Meta: meta,
    })
	// penggunaan oklist(c, "isiMessage", namaVarData, namaVarMeta)
}

func created(c *fiber.Ctx, message string, data any, location string) error { // set success true, menampilkan message dan data dari parameter webresponse 
    c.Set("Location", location) // memberi tahu klien di mana sumber daya baru berada
    return c.Status(fiber.StatusCreated).JSON(WebResponse{ 
        Success: true, Message: message, Data: data, // set status = 201 post berhasil
    })
}

func noContent(c *fiber.Ctx) error { 
    return c.SendStatus(fiber.StatusNoContent) // set : 204 delete/put berhasil
	// penggunaan noContent(c) tanpa body karena ga ngelakuin apa apa selain set status
}
 
func fail(c *fiber.Ctx, status int, message string) error { // set success, dan menampilkan message sesuai value parameter web response
    return c.Status(status).JSON(WebResponse{Success: false, Message: message}) //set status melalui parameter bukan set fixed status, karena fail ada berbagai macam jadi dideklarasikan melalui parameter saat fungsi fail digunakan.
}
 
func failValidation(c *fiber.Ctx, errs map[string]string) error {
    return c.Status(fiber.StatusUnprocessableEntity).JSON(WebResponse{ // set status : 422 eror karena bagian data tidak sesuai permintaan dan menampilkan pesan eror
        Success: false, Message: "validasi gagal", Errors: errs,
    })
	// buat pesan validasi 
	// penggunaan failValidation("namaField" : "isiMessage", dst)
}

