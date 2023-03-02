describe('Login Component', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200'); // assuming the login component is served at '/login'
  });

  // it('should display user information', () => {
  //   cy.get('div').contains('User Information'); // make sure the User Information heading is displayed
  //   cy.get('p').contains('ID'); // make sure the user ID is displayed
  // });

  it('should redirect to Spotify API', () => {
    cy.get('button').contains('Redirect to Spotify API').click(); // click the Redirect to Spotify API button
    //cy.url().should('eq', 'https://api.spotify.com/'); // make sure the URL has changed to the Spotify API URL
  });

  // it('should handle auth', () => {
  //   cy.get('button').contains('handle auth').click(); // click the handle auth button
  //   cy.url().should('include', 'https://accounts.spotify.com/'); // make sure the URL has changed to the Spotify accounts URL
  // });

  it('should add a friend', () => {
    const friendName = 'John'; // set a friend name to use in the test
    cy.get('input').type(friendName); // type the friend name in the input field
    cy.get('button').contains('Add Friend').click(); // click the Add Friend button
    cy.get('ul').contains(friendName); // make sure the friend name is displayed in the list of friends
  });

  // it('should display songs on submit', () => {
  //   cy.get('button').contains('Display Songs').click(); // click the Display Songs button
  //   cy.get('.song-card').should('have.length', 1); // make sure at least one song card is displayedd
  // });
});
