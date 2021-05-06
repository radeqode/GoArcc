import React, { useContext, useEffect } from 'react'
import { useRouter } from 'next/router';
import { UserContext } from '../../../Contexts/UserContext';
import { SERVER } from '../../../utils/constants';

function Callback() {
    const router = useRouter();
    const { setUser, user } = useContext(UserContext)

    const { query } = router
    useEffect(() => {
        if (query.code && query.code.length > 0) {
            (async () => {
                console.log(query)
                if (query.code && user.userId != "") {


                    let res = await fetch(`${SERVER}/account/get-user-all-account/${user.userId}`, {
                        headers: new Headers({
                            'Authorization': `Bearer ${user.idToken}`,
                        })
                    })
                    let data = await res.json()
                    const newUser = user

                    newUser.accounts = data.accounts || []

                    // todo: gql
                    res = await fetch(`${SERVER}/vcs-connection/GITHUB/callback?code=${query.code}&account_id=${newUser.accounts[0].id}`, {
                        headers: new Headers({
                            'Authorization': `Bearer ${user.idToken}`,
                        })
                    })
                    data = await res.json()
                    newUser.provider = data.provider

                    console.log(newUser)
                    setUser(newUser)
                    router.push("/tell-us-more")
                } else {
                    router.push("/")
                }

            })()

        }
        // else
        // router.push("/")
    }, [query, router, setUser, user])

    return (
        <div>

        </div>
    )
}

export default Callback

