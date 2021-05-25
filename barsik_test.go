package main

import (
       "testing"
)

func TestConvert(t *testing.T) {
     convertLinksToS3("https://barsik-geometry-1-images.s3-us-west-2.amazonaws.com/",
		      "public/index.html",
		      "src")

}