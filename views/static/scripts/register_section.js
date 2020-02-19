class RegisterSection {
    constructor() {
        this.form = document.getElementById("register_form");

        this.form.onsubmit = this.on_submit.bind(this);
    }

    async on_submit(e){
        e.preventDefault();
        var data = {
            name:this.form.children[0].value,
            password:this.form.children[1].value,
            confirm_password:this.form.children[2].value
        };
        var options = {
            method:'POST',
            body:JSON.stringify(data)
        };
        var response = await fetch("/user", options);
        var data = await response.json();
        console.log(data);


    }

}

var brain = new RegisterSection();