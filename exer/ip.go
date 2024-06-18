import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

// downloadFile downloads an object from GCS.
func downloadFile(w io.Writer, bucket, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	if _, err := io.Copy(w, rc); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	fmt.Fprintf(w, "Blob %v downloaded.\n", object)
	return nil
}
  
