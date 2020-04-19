# Get all of the assignee's candidates

Send a message from one user to another.

**URL** : `/assignee/findAssigneesCandidates/{_id}`

**Method** : `GET`

**Data constraints**

Provide the name of the assignee in the url.

**Example** 

```
 /assignee/findAssigneesCandidates/5e9c4a78323cc69c4f34da78`
```

## Success Response

**Condition** : If everything is OK and the assignee exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Assignee's ID": "5e9b2e37a358a83e95401984"
}

```

## Error Response

**Condition** : If the assignee doesn't exist in the database.

**Code** : `404 NOT FOUND`

**Content example** :

```json
{
    "error": "Assignee with the given name doesnot exist"
}
```
