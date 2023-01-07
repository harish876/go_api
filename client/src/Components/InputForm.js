import React, { useState,useRef, useEffect } from 'react';
import { Button, Checkbox, Form, Input,Select } from 'antd';
import { isEmpty } from 'lodash';
import axios from 'axios'
const {Item} = Form
const InputForm = (props) => {
 
    const {data,type,refreshPage} = props
    const formRef= useRef()
    useEffect(()=>{
      formRef.current.resetFields();
    },[data])
    const postData =async(values)=>{
        values.status = values.status==='true'?true:false
        if(!isEmpty(type) && ['create'].includes(type))
        {
            await axios.post('http://127.0.0.1:4000/api/users/create',values)
        }
        else if(!isEmpty(type) && ['update'].includes(type))
        {
            await axios.patch(`http://127.0.0.1:4000/api/users/update/${data._id}`,values)
        }
        
    }
    const onFinish = (values) => {

        postData(values)
        console.log('Success:', values);
        refreshPage()
    };
    const onFinishFailed = (errorInfo) => {
        console.log('Failed:', errorInfo);
    };
  return (
    <Form
      ref={formRef}
      name="basic"
      labelCol={{span: 8,}}
      wrapperCol={{span: 16,}}
      initialValues={{...data}}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="on"
    >
      <Item
        label="Name"
        name="name"
        rules={[
          {
            required: true,
            message: 'Please input your username!',
          },
        ]}
      >
        <Input />
      </Item>
        
      <Item
        label="Email"
        name="email"
        rules={[
          {
            required: true,
            message: 'Please input your email!',
          },
        ]}
      >
        <Input />
      </Item>

      <Item
        label="Password"
        name="password"
        rules={[
          {
            required: true,
            message: 'Please input your password!',
          },
        ]}
      >
        <Input.Password />
      </Item>

        <Item label="Select" name="status">
          <Select>
            <Select.Option value="true">Active</Select.Option>
            <Select.Option value="false">Inactive</Select.Option>
          </Select>
        </Item>

      <Item
        wrapperCol={{
          offset: 8,
          span: 16,
        }}
      >
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Item>
      </Form>
  );
};
export default InputForm;