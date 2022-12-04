## Golang Microservice

**Problem**

To Implement two endpoints using golang as core;
1. For the first task Your job to implement /hai/ endpoint that accepts any of the three parameters (id  - hai - city) and return hai data in JSON format. You have to query the database to get the results. Additionally, string parameters (hai-city) may contain part or full expected string, and you have to search the database for a possible list of matches.
2. For the second task Your job to implement /haicounter/ endpoint that returns list of cities and number of hais in each city. You have to query the database to get the results.


**Implementation**

I have solved the given problem using a five-layer approach as following;

![Five_layer Model](https://i.imgur.com/3Z0XsOA.png)

**api_call->endpoint->handler->service->database->Model**

This approach helps in mainting clean code along with unit-tests at different layers. For the given task, I have implemented unit-testing at main-body and database layer. Following are the unit-test execution snapshots;

#### **Database Unit Testing**

![database_unit-Testing](https://i.imgur.com/kyRfokg.png)

#### **Endpoints Unit Testing**

![endpoints_unit_testing](https://i.imgur.com/UP3ymAU.png)


#### **Use Cases**
Considering the first endpoint ***/hai/***, all the mentioned use-cases are covered that are;
1. To show result for id parameter
2. To show result for city parameter
3. To show result for hai parameter
4. To show result for city and hai paramter

Adding some snapshots for given endpoints calls;

### 1. **localhost:3333/haicounter/**

![haicounter](https://i.imgur.com/WSbuRDB.png)


### 2. **localhost:3333/hai/?id=4**

![hai-endpoint](https://i.imgur.com/MWD14xZ.png)

### 3. **localhost:3333/hai/?hai=م

![hai-hai-endpoint](https://i.imgur.com/pt40RFO.png)

### 4. **localhost:3333/haicounter/ (Method:POST)**

![method-restriction](https://i.imgur.com/OiWEXLd.png)
---


### 5. **localhost:3333/simsim (checking for path restriction)**

![wrong-endpoint](https://i.imgur.com/Kj3aljj.png)



The microserice can be further improved as follows;

1. golangci-lint for code standarization check
2. Make file for automation
3. Runtime DB instantance initialization
4. adding logs
5. Dockerization

