# Integration Test Cases

## Order Management

- **createOrder:**
  - [x] Test creating an order with valid data.
  - [x] Test creating an order with invalid data.

- **updateOrder:**
  - [x] Test updating an order with valid data.
  - [x] Test updating an order without sufficient permissions.

- **deleteOrder:**
  - [x] Test deleting an order with valid ID.
  - [x] Test deleting an order without sufficient permissions.

## Contract Management

- **createContract:**
  - [x] Test creating a contract with valid data.
  - [x] Test creating a contract without sufficient permissions.
  - [x] Test creating a contract with invalid data.

- **updateContract:**
  - [x] Test updating a contract with valid data.
  - [x] Test updating a contract without sufficient permissions.
  - [x] Test updating a non-existent contract.

- **deleteContract:**
  - [x] Test deleting a contract with valid ID.
  - [x] Test deleting a contract without sufficient permissions.
  - [x] Test deleting a non-existent contract.

## Bid Management

- **createBid:**
  - [x] Test creating a bid with valid data.
  - [x] Test create bid when auction is not ongoing.

## Station Management

- **createStation:**
  - [x] Test creating a station with valid data.
  - [x] Test creating a station without sufficient permissions.
  - [x] Test creating a station with invalid data.

- **updateStation:**
  - [x] Test updating a station with valid data.
  - [x] Test updating a station without sufficient permissions.
  - [x] Test updating a non-existent station.

- **deleteStation:**
  - [x] Test deleting a station with valid ID.
  - [x] Test deleting a station without sufficient permissions.
  - [x] Test deleting a non-existent station.

- **offSetStationConsumption:**
  - [ ] Test offsetting station consumption with valid data.
  - [ ] Test offsetting station consumption with invalid data.

## Auction Management

- **createAuction:**
  - [x] Test creating an auction with valid data.
  - [x] Test creating an auction without sufficient permissions.
  - [x] Test creating an auction with invalid data.

- **updateAuction:**
  - [x] Test updating an auction without sufficient permissions.
  - [x] Test updating a non-existent auction.

- **finishAuction:**
  - [ ] Test finishing an auction that has expired.
  - [ ] Test finishing an auction without sufficient permissions.
  - [ ] Test finishing an auction that hasn't expired yet.

## User Management

- **withdrawStablecoin:**
  - [x] Test withdrawing funds with valid data.
  - [x] Test withdrawing funds with insufficient balance.

- **withdrawVolt:**
  - [x] Test withdrawing funds with valid data.
  - [x] Test withdrawing funds with insufficient balance.

- **createUser:**
  - [x] Test creating a user with valid data.
  - [x] Test creating a user without sufficient permissions.
  - [x] Test creating a user with invalid data.

- **updateUser:**
  - [x] Test updating a user with valid data.
  - [x] Test updating a user without sufficient permissions.
  - [x] Test updating a non-existent user.

- **withdrawApp:**
  - [ ] Test withdrawing application funds with valid data.
  - [ ] Test withdrawing application funds without sufficient permissions.

- **deleteUser:**
  - [x] Test deleting a user with valid address.
  - [x] Test deleting a user without sufficient permissions.
  - [x] Test deleting a non-existent user.
