import { parseJwt } from "@/utils/jwtUtils";

export const fetchBooksAll = async (type) => {
    const token = localStorage.getItem('token');
    if (!token) return { error: 'No token found' };

    const { user_id } = parseJwt(token);
    const headers = {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
    };

    try {
        const response = await fetch('http://localhost:8080/books');
        const allBooks = await response.json();

        const readResponse = await fetch(`http://localhost:8080/users/${user_id}/read-books`, {
            method: 'GET',
            headers
        });
        const readBooks = await readResponse.json();

        const abandonedResponse = await fetch(`http://localhost:8080/users/${user_id}/abandoned-books`, {
            method: 'GET',
            headers
        });
        const abandonedBooks = await abandonedResponse.json();

        return allBooks
            .filter(book => book.type === type)
            .map(book => {
                const isRead = readBooks.some(readBook => readBook.ID === book.ID);
                const isAbandoned = abandonedBooks.some(abandonedBook => abandonedBook.ID === book.ID);
                return { ...book, read: isRead, abandoned: isAbandoned };
            });
    } catch (error) {
        console.error('Ошибка:', error);
        return { error };
    }
};
