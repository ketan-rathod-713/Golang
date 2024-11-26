import React from 'react';

const Pagination = ({ page, limit, total, handlePageChange }) => {
  const totalPages = Math.ceil(total / limit);
  const pagesToShow = 10; // Number of pages to show in the pagination bar

  const renderPagination = () => {
    const currentPage = page || 1;
    const startPage = Math.max(1, currentPage - Math.floor(pagesToShow / 2));
    const endPage = Math.min(totalPages, startPage + pagesToShow - 1);

    let pages = [];
    for (let i = startPage; i <= endPage; i++) {
      pages.push(i);
    }

    return (
      <>
        {currentPage > 1 && (
          <button className='page-btn' onClick={() => handlePageChange(currentPage - 1)}>Prev</button>
        )}
        {pages.map((pageNum) => (
          <button
            key={pageNum}
            onClick={() => handlePageChange(pageNum)}
            className={`${pageNum === currentPage ? "current-page page-btn" : "page-btn"}`}
            style={{ fontWeight: pageNum === currentPage ? 'bold' : 'normal' }}
          >
            {pageNum}
          </button>
        ))}
        {currentPage < totalPages && (
          <button className='page-btn' onClick={() => handlePageChange(currentPage + 1)}>Next</button>
        )}
      </>
    );
  };

  return <div className='pagination'>{totalPages > 1 && renderPagination()}</div>;
};

export default Pagination;
