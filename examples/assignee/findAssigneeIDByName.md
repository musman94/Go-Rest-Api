# Get the assignee's ID by name

Send a message from one user to another.

**URL** : `/assignee/findAssigneeIDByName{name}`

**Method** : `GET`

**Data constraints**

Provide the name of the assignee in the url.


**Example** 

```
/assignee/findAssigneeIDByName/ali`
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
