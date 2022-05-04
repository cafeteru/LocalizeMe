// login.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('Login user', () => {
    it('should login with correct data', () => {
        const email = 'admin@email.es';
        cy.visit('http://localhost:4200');
        const showLogin = cy.get('#showLogin');
        showLogin.click();
        cy.get('#loginEmail').type(email).should('have.value', email);
        cy.get('#loginPassword').type(email).should('have.value', email);
        cy.get('#loginEnter').click();
        showLogin.should('not.exist');
    });

    it('should login with incorrect data', () => {
        cy.visit('http://localhost:4200');
        const showLogin = cy.get('#showLogin');
        showLogin.click();
        cy.get('#loginEmail').type('admin@email.es').should('have.value', 'admin@email.es');
        const loginEnter = cy.get('#loginEnter');
        loginEnter.should('be.disabled');
        cy.get('#loginPassword').type('incorrect').should('have.value', 'incorrect');
        loginEnter.click();
        showLogin.should('exist');
    });
});
