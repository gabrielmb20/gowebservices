package main

import (
    "context"
    pb "github.com/gabrielmb20/gowebservices/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
)

func main() {
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    // Agregar libro
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Operating System Concepts",
        Edition:   "9th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("\n\nNo se pudo agregar el libro: %v", err)
    }

    // Obtener libro
    log.Printf("\n\nLibro creado ID: %s", r.Value)
    book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("\n\nEl libro consultado no existe: %v", err)
    }
    log.Printf("\n\nLibro consultado: ", book.String())

    // Eliminar libro
    bookDel, err := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("\n\nNo se pudo eliminar el libro: %v", err)
    }
    log.Printf("\n\nLibro eliminado: ", bookDel.String())

    // Obtener libro nuevamente para comprobar si se elimino
    bookGet, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("\n\nEl libro consultado no existe: %v", err)
    }
    log.Printf("\n\nLibro consultado: ", bookGet.String())

    // Actualizar libro
    upd, err := c.UpdateBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "UPDATED Book",
        Edition:   "10th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("\n\nNo se pudo actualizar el libro: %v", err)
    }

    // Obtener libro nuevamente para comprobar si se actualizo
    bookGet, err := c.GetBook(ctx, &pb.BookID{Value: upd.Value})
    if err != nil {
        log.Fatalf("\n\nEl libro consultado no existe: %v", err)
    }
    log.Printf("\n\nLibro consultado: ", bookGet.String())
}
