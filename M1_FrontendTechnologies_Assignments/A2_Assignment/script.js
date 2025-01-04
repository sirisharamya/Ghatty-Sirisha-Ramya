const expenseForm = document.getElementById('expenseForm');
const expenseTable = document.getElementById('expenseTable');
const categorySummary = document.getElementById('categorySummary');
let expenses = JSON.parse(localStorage.getItem('expenses')) || [];
function renderExpenses() {
  expenseTable.innerHTML = '';
  categorySummary.innerHTML = '';
  expenses.forEach((expense, index) => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>$${expense.amount}</td>
      <td>${expense.description}</td>
      <td><span class="badge bg-primary">${expense.category}</span></td>
      <td>
        <button class="btn btn-danger btn-sm" onclick="deleteExpense(${index})">
          <i class="bi bi-trash"></i> Delete
        </button>
      </td>
    `;
    expenseTable.appendChild(row);
  });
  const categoryTotals = {};
  expenses.forEach(expense => {
    categoryTotals[expense.category] = (categoryTotals[expense.category] || 0) + expense.amount;
  });

  for (const [category, total] of Object.entries(categoryTotals)) {
    const listItem = document.createElement('li');
    listItem.className = 'list-group-item d-flex justify-content-between align-items-center';
    listItem.innerHTML = `
      ${category}
      <span class="badge bg-success">$${total}</span>
    `;
    categorySummary.appendChild(listItem);
  }
  localStorage.setItem('expenses', JSON.stringify(expenses));
}
expenseForm.addEventListener('submit', (e) => {
  e.preventDefault();

  const amount = parseFloat(document.getElementById('amount').value);
  const description = document.getElementById('description').value.trim();
  const category = document.getElementById('category').value;

  if (amount && description && category) {
    expenses.push({ amount, description, category });
    expenseForm.reset();
    renderExpenses();
  }
});
function deleteExpense(index) {
  expenses.splice(index, 1);
  renderExpenses();
}
renderExpenses();
