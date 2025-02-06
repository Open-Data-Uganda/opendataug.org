import { z } from 'zod';

export const LoginSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim().toLowerCase(),
  password: z.string().min(1, 'Password is required')
});

export const ResetPasswordSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim().toLowerCase()
});

export const SignUpSchema = z.object({
  firstName: z.string().min(1, 'First name is required').trim(),
  lastName: z.string().min(1, 'Last name is required').trim(),
  email: z.string().email('Invalid email address').trim().toLowerCase()
});

export const APIKeySchema = z.object({
  name: z.string().min(1, 'API Key name is required').max(10, 'API Key name cannot exceed 10 characters').trim()
});

export const InviteUserSchema = z.object({
  email: z.string().min(1, 'Email is required').email().trim(),
  first_name: z.string().min(1, 'First name is required'),
  other_name: z.string().min(1, 'Last name is required'),
  telephone: z.string().optional(),
  role: z.string().min(1, 'Role is required')
});

export const SetPasswordSchema = z
  .object({
    password: z
      .string()
      .min(1, 'Password is required')
      .min(8, 'Password must be at least 8 characters long')
      .regex(/^(?=.*[0-9])(?=.*[!@#$%^&*])/, 'Password must contain at least one number and one special character'),
    confirm_password: z.string().min(1, 'Confirm password is required')
  })
  .refine((data) => data.password === data.confirm_password, {
    message: "Passwords don't match",
    path: ['confirm_password']
  });

export const EditProfileSchema = z.object({
  first_name: z.string().optional(),
  other_name: z.string().optional(),
  email: z.string().optional(),
  phone: z.string().optional()
});
