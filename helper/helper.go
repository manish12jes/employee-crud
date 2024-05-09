package helper

// Pagination Object
type Paginate struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func Pagination(limit, page, recordcount int) *Paginate {
	pagination := Paginate{}
	total := (recordcount / limit)

	// Calculator Total Page
	remainder := (recordcount % limit)
	if remainder == 0 {
		pagination.TotalPage = total
	} else {
		pagination.TotalPage = total + 1
	}

	pagination.CurrentPage = page
	pagination.RecordPerPage = limit

	if page <= 0 {
		pagination.Next = page + 1
	} else if page < pagination.TotalPage {
		pagination.Previous = page - 1
		pagination.Next = page + 1
	} else if page == pagination.TotalPage {
		pagination.Previous = page - 1
		pagination.Next = 0
	}

	return &pagination
}
