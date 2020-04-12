package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	"../blogpb"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Mark",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected err: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// read Blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5e93411f11844eecb9e7787d"})
	if err2 != nil {
		fmt.Printf("Error happened whilst reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened whilst reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happened whilst updating: %v\n", updateErr)
	}
	fmt.Printf("Blog was read: %v\n", updateRes)
}
