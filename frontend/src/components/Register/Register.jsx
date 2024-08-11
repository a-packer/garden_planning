import React, {useState} from 'react'

const Register = () => {

  const [message, setMessage] = useState('');
  const [registerUsername, setRegisterUsername] = useState('');
  const [registerPassword, setRegisterPassword] = useState('');

  const handleSubmitRegister = async (e) => {
    e.preventDefault();
    console.log('register', registerUsername, registerPassword)
    try {
      const response = await fetch('/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ registerUsername, registerPassword }),
      });
      if (response.ok) {
        setMessage("Registration complete");
      } else {
        setMessage('Invalid username or password for registration.');
      }
    } catch (error) {
      setMessage('Error registering');
    }
  };

  return (
    <div className="registerWrapper">
      <h2>Register New User</h2>
      <form onSubmit={handleSubmitRegister}>
          <div>
              <label>Username:</label>
              <input type="text" value={registerUsername} onChange={(e) => setRegisterUsername(e.target.value)} required />
          </div>
          <div>
              <label>Password:</label>
              <input type="password" value={registerPassword} onChange={(e) => setRegisterPassword(e.target.value)} required />
          </div>
          <button type="submit">Register</button>
      </form>
      {message && <p>{message}</p>}
    </div>
  )
}

export default Register