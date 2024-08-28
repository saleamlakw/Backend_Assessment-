package usecases

import (
	"context"

	"github.com/saleamlakw/LoanTracker/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookUserCase struct {
	bookRepository entities.BookRepositary
    logRepo entities.LogRepository
}

func NewBookUserCase(bookRepository entities.BookRepositary, logRepo entities.LogRepository) *bookUserCase {
	return &bookUserCase{
		bookRepository: bookRepository,
        logRepo: logRepo, 
	}
}

func (bu *bookUserCase) ApplyBorrowRequest(ctx context.Context,bookID, userID primitive.ObjectID) error {
    borrowRequest := entities.BorrowRequest{
        BookID: bookID,
        UserID: userID,
        Status: "pending",
    }
    _, err := bu.bookRepository.ApplyBorrowRequest(ctx,borrowRequest)
    if err!=nil{
        return err
    }
    // Log the book borrow request submission
    err = bu.logRepo.LogEvent(ctx,"Book Borrow Request Submission", "User requested to borrow a book.")
    return err
}

func (bu *bookUserCase) GetBorrowRequestByID(ctx context.Context,id primitive.ObjectID) (entities.BorrowRequest, error) {
    return bu.bookRepository.GetBorrowRequestByID(ctx,id)
}

func (bu *bookUserCase) GetAllBorrowRequests(ctx context.Context,status, order string) ([]entities.BorrowRequest, error) {
    return bu.bookRepository.GetAllBorrowRequests(ctx,status, order)
}

func (bu *bookUserCase) UpdateBorrowRequestStatus(ctx context.Context,id primitive.ObjectID, status string) error {
       err := bu.bookRepository.UpdateBorrowRequestStatus(ctx,id, status)

       if err != nil {
           return err
       }

       // Log the status update
       event := "Book Borrowing Status Update"
       details := "Borrow request status updated to " + status
       err = bu.logRepo.LogEvent(ctx,event, details)
       return err
}

func (bu *bookUserCase) DeleteBorrowRequest(ctx context.Context,id primitive.ObjectID) error {
    return bu.bookRepository.DeleteBorrowRequest(ctx,id)
}
