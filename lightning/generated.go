// BookService exposes methods for querying of the library.
type BookServiceClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb.CallOption) (*Book, error)
	
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb.CallOption) (BookService_QueryBooksClient, error)
	
	MakeCollection(ctx context.Context, opts ...grpcweb.CallOption) (BookService_MakeCollectionClient, error)

	BookChat(ctx context.Context, opts ...grpcweb.CallOption) (BookService_BookChatClient, error)
}

type bookServiceClient struct {
	client *grpcweb.Client
}

// NewBookServiceClient creates a new client to a BookService running on the hostname.
func NewBookServiceClient(hostname string, opts ...grpcweb.DialOption) BookServiceClient {
	return &bookServiceClient{
		client: grpcweb.NewClient(hostname, "library.BookService", opts...),
	}
}
