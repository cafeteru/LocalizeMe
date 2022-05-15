import { login } from '../utils.spec';

describe('Update me', () => {
    let email;

    beforeEach(() => {
        email = Cypress.env('email');
    });

    it('should update owner information of a user', () => {
        login();
        cy.get('#appUserInfo').click();
        cy.get('#updateMeButton').click();
        const randomNumber = new Date().getTime();
        const email = `user${randomNumber}@email.es`;
        Cypress.env('email', email);
        const updateUserEmail = cy.get('#updateUserEmail');
        updateUserEmail.clear().type(email).should('have.value', email);
        cy.get('#updateUserChangePasswordYes').click();
        cy.get('#updateUserPassword').type(email).should('have.value', email);
        cy.get('#updateUserCheckPassword').type(email).should('have.value', email);
        cy.get('#updateUserSend').click();
        updateUserEmail.should('not.exist');
    });
});
